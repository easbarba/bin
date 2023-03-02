#!/usr/bin/env php
<?php
declare(strict_types=1);

namespace bin;

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
require $home_dir . "/.config/composer/vendor/autoload.php";

// API
final class Response
{
    const URL = "https://gist.githubusercontent.com/funkyhippo/1d40bd5dae11e03a6af20e5a9a030d81/raw/?";

    public function __construct(public string $api_file_cached)
    {
    }

    private function guzzly(): string
    {
        $client = new GuzzleHttp\Client();
        $response = $client->get(self::URL);

        return $response->getBody()->getContents();
    }

    private function curly(): array
    {
        $handler = curl_init();
        curl_setopt_array($handler, [
            CURLOPT_URL => self::URL,
            CURLOPT_HTTPGET => true,
            CURLOPT_RETURNTRANSFER => true,
        ]);

        $response_json = curl_exec($handler);
        curl_close($handler);

        return $response_json;
    }

    private function write(string $grabbed): void
    {
        file_put_contents($this->api_file_cached, $grabbed);
    }

    public function chapters(): array
    {
        if (file_exists($this->api_file_cached)) {
            $fool = file_get_contents($this->api_file_cached);
            $meh = json_decode($fool, true);
            return $meh["chapters"];
        }

        $grabbed = $this->guzzly();
        $this->write($grabbed);

        return json_decode($grabbed, true)["chapters"];
    }
}

// CLI
$opm_dir = $home_dir . "/Downloads/One Punch Man";
$destination_dir = $argv[1] ?? $opm_dir; // if no directory has been set, use default.
$destination_dir = realpath($destination_dir);
if (!file_exists($destination_dir)) {
    mkdir($destination_dir);
}

$api_file = $destination_dir . "/opm.json";
// ACTION

// $chapters_file = $home_download_dir
$chapters = (new Response($api_file))->chapters();
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
