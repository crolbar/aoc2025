<?php

$COOKIE = getenv("AOC_COOKIE");
const URL = "https://adventofcode.com/2025/day/__NUMDAY__/input";

if (count($argv) <= 1 || !$argv[1]) {
    exit("no arg2\n");
}

function getFmtDay(string $str): string {
    $day=$str;

    $fmt = new NumberFormatter("en", NumberFormatter::SPELLOUT);

    if (!str_contains($day, "day")) {
        if (ctype_digit($day)) {
            // $day="day"
            $d=ucfirst($fmt->format($day));
            $day="day$d";
        } else {
            $d=ucfirst($day);
            $day="day$d";
        }
    }

    return $day;
}

if ($argv[1] == "new") {
    if (count($argv) < 3) {
        exit("no arg 3");
    }

    if (!ctype_digit($argv[2])) {
        exit("use digit for the day, ex: 1, 2\n");
    }

    $name=getFmtDay($argv[2]);


    $b = touch($name . ".go");
    if (!$b) {
        exit("faled to create .go file");
    }

    $url = str_replace("__NUMDAY__", $argv[2], URL);
    echo "fetching input: " .$url . "\n";

    $conts = file_get_contents($url, false, stream_context_create([
        "http" => [ "method"  => "GET", "header"  => "Cookie: $COOKIE;" ]
    ]));

    file_put_contents("$name.input", $conts);

    echo "input is in $name.input\n";
    exit();
}


$name = getFmtDay($argv[1]);
exec("go run $name.go $name.input");
