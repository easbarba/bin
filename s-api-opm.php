#!/usr/bin/php
<?php

declare(strict_types=1);

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

$home = getenv('HOME');
$homeDown = $home . "/Downloads/One Punch Man";
$to = $homeDown;

// if no directory has been set, use default.
if(!empty($argv[1])) {
    $to = $argv[1];

    if(!file_exists($to)) {
        mkdir($to);
    }
}

// API
function guzzly(string $homedir): array
{
    require($homedir . '/.config/composer/vendor/autoload.php');

    $url = 'https://gist.githubusercontent.com/funkyhippo/1d40bd5dae11e03a6af20e5a9a030d81/raw/?';
    $client = new GuzzleHttp\Client();
    $response = $client->get($url);

    return json_decode($response->getBody()->getContents(), true);
}

function curly(string $url): array
{
    $handler = curl_init();
    curl_setopt_array($handler, [
        CURLOPT_URL => $url,
        CURLOPT_HTTPGET => true,
        CURLOPT_RETURNTRANSFER => true
    ]);

    $response_json = curl_exec($handler);
    curl_close($handler);

    return json_decode($response_json, true);
}


// ACTION
$chapters = guzzly($home)['chapters'];
foreach ($chapters as $chapter) {
    $title = "{$chapter["title"]}";
    $fld = $to . DIRECTORY_SEPARATOR . $title;

    // create dir of chapter if needed
    if(!file_exists($fld) && !is_dir($fld)) {
        echo "creating directory: {$fld}";
        mkdir($fld, 0755);
    }

    // select pictures array
    $pics = $chapter['groups']["/r/OnePunchMan"];

    // Check if folder has at least one pic
    // so to avoid API pics naming change.
    if((count(scandir($fld)) <= 2)) {
        foreach ($pics as $pic) {
            $img =  $fld . DIRECTORY_SEPARATOR . basename($pic);

            // download file only if needed
            if(!file_exists($img)) {
                echo "\nGetting pic: {$img}";
                file_put_contents($img, file_get_contents($pic));
            }
        }
    }
}

// echo $title, " api: ", count($pics), " folder: ", (count(scandir($fld)) - 2), "\n";
