#!/usr/bin/env php
<?php
declare(strict_types=1);

namespace bin;

use Exception;
use GuzzleHttp\Client;
use GuzzleHttp\Exception\ClientException;
use GuzzleHttp\Psr7;

/*
 * Bin is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Bin is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Bin. If not, see <https://www.gnu.org/licenses/>.
 */

# TODO: cache api response to a file for future requests.

$home_dir = getenv("HOME");
require $home_dir .
    DIRECTORY_SEPARATOR .
    ".config/composer/vendor/autoload.php";

// API
final class Response
{
    public function __construct(
        private string $url,
        private string $config_file
    ) {
    }

    // use php's stream-handling
    private function native(): string
    {
        $context = stream_context_create([
            "http" => [
                "method" => "GET",
                "header" => [
                    "Accept: application/json",
                    "Content-Type: application/json",
                ],
            ],
        ]);

        $response_json = file_get_contents($this->url, false, $context);

        return $response_json;
    }

    private function guzzly(): string
    {
        try {
            $client = new Client(["base_uri" => $this->url]);
            $response = $client->get("");

            return $response->getBody()->getContents();
        } catch (ClientException $e) {
            echo Psr7\Message::toString($e->getRequest());
            echo Psr7\Message::toString($e->getStatusCode());
            exit();
        }
    }

    private function curly(): array
    {
        $handler = curl_init();
        curl_setopt_array($handler, [
            CURLOPT_URL => $this->url,
            CURLOPT_HTTPGET => true,
            CURLOPT_HTTPHEADER => ["Accept: application/json"],
            CURLOPT_RETURNTRANSFER => true,
        ]);

        $response_json = curl_exec($handler);
        curl_close($handler);

        return $response_json;
    }

    // cache response contents to a file
    private function cache(string $grabbed): void
    {
        file_put_contents($this->config_file, $grabbed);
    }

    public function chapters(): ?array
    {
        if (file_exists($this->config_file)) {
            $config_content = file_get_contents($this->config_file);
            $response = json_decode($config_content, true);

            return $response["chapters"];
        }

        $grabbed = $this->native();
        if (is_null($grabbed)) {
            $this->cache($grabbed);
        } else {
        }

        return json_decode($grabbed, true)["chapters"];
    }
}

// CLI
// TODO: add a proper cli options with --force, --cache, --url URL, --to DIR
$destination_dir = "";
$default_folder =
    $home_dir .
    DIRECTORY_SEPARATOR .
    "Downloads" .
    DIRECTORY_SEPARATOR .
    "One Punch Man";

$destination_dir = realpath(count($argv) === 1 ? $default_folder : $argv[1]);
if (!file_exists($destination_dir)) {
    mkdir($destination_dir);
}

$config_file = $destination_dir . "/opm.json";

echo "destination: ", $destination_dir, "\n";
echo "config file: ", $config_file, "\n";

# BEGIN
const URL = "https://gist.githubusercontent.com/funkyhippo/1d40bd5dae11e03a6af20e5a9a030d81/raw/?";
$chapters = (new Response(URL, $config_file))->chapters();

// ACTION
foreach ($chapters as $chapter) {
    $chapter_title = "{$chapter["title"]}";
    $chapter_dir = $destination_dir . DIRECTORY_SEPARATOR . $chapter_title;

    // create dir of chapter if needed
    if (!file_exists($chapter_dir) && !is_dir($chapter_dir)) {
        echo "creating directory: {$chapter_dir}";
        mkdir($chapter_dir, 0755);
    }

    // select pictures array
    $pictures = $chapter["groups"]["/r/OnePunchMan"];

    // Check if folder has at least one pic
    // so to avoid API pics naming change.
    if (count(scandir($chapter_dir)) <= 2) {
        foreach ($pictures as $picture) {
            $pic_name_decoded = base64_decode(basename($picture));
            $pic_real_name = substr(
                $pic_name_decoded,
                strpos($pic_name_decoded, "_") + 1
            );
            $pic_path = $chapter_dir . DIRECTORY_SEPARATOR . $pic_real_name;

            // download file only if needed
            if (!file_exists($pic_path)) {
                echo "\nGetting pic: {$pic_path}";
                file_put_contents($pic_path, file_get_contents($picture));
            }
        }
    } else {
        echo "DONE: {$chapter_title}\n";
    }
}
