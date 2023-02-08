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

$homeDown = getenv('HOME') . DIRECTORY_SEPARATOR . "Downloads" . DIRECTORY_SEPARATOR . "One Punch Man";
$to = $argv[1] ?? $homeDown;
if(!file_exists($to)) mkdir($to);

// API

$url = 'https://gist.githubusercontent.com/funkyhippo/1d40bd5dae11e03a6af20e5a9a030d81/raw/?';
$ch = curl_init($url);
curl_setopt($ch, CURLOPT_HTTPGET, true);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
$response_json = curl_exec($ch);
curl_close($ch);
$response=json_decode($response_json, true);

// GRABBING VOLUMES
$chapters = $response['chapters'];
foreach ($chapters as $chapter) {
    $title = "{$chapter["title"]}";
    $fld = $to . DIRECTORY_SEPARATOR . $title;

    if(!file_exists($fld) && !is_dir($fld)) {
        echo "creating directory: {$fld}";
        mkdir($fld, 0755);
    }

    $pics = $chapter['groups']["/r/OnePunchMan"];
    foreach ($pics as $pic) {
        $img =  $fld . DIRECTORY_SEPARATOR . basename($pic);
        if(!file_exists($img)) {
            echo "\nGetting pic: {$img}";
            file_put_contents($img, file_get_contents($pic));
        }
    }
}
