<?php


error_reporting(E_ALL ^ E_WARNING);

function get_input() {
  return preg_split("/,/", preg_split("/\n/", file_get_contents("./input.txt"))[0]);
}

function compute($items, $noun, $verb) {
  $i = 0;
  // as requested:
  $items[1] = $noun;
  $items[2] = $verb;
  for (;;) {
    switch ($items[$i]) {
      case 1:
        $items[$items[$i+3]] = $items[$items[$i+1]] + $items[$items[$i+2]];
        break;
      case 2:
        $items[$items[$i+3]] = $items[$items[$i+1]] * $items[$items[$i+2]];
        break;
      case 99:
        break 2;
    }
    $i += 4;
    if ($i > count($items)) {
      break;
    }
  }
  return $items[0];
}

function p1($items) {
  return compute($items, 12, 2);
}

function p2($items) {
  // need to find: 19690720
  for ($i=0;$i<100;$i++) {
    for ($j=0;$j<100;$j++) {
      if (compute($items, $i, $j) == 19690720) {
        return 100 * $i + $j;
      }
    }
  }
}

// 4484226
echo p1(get_input());

echo "\n";

// 5696
echo p2(get_input());
