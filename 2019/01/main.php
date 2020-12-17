<?php
error_reporting(E_ALL ^ E_WARNING);

function get_input() {
  return preg_split("/\n/", file_get_contents("./input.txt"));
}

function p1($items) {
  $ret = 0;
  foreach ($items as $item) {
    if ($item == "") {
      break;
    }
    $ret += floor($item / 3) - 2;
  }
  return $ret;
}

function p2($items) {
  $ret = 0;
  foreach ($items as $item) {
    if ($item == "") {
      break;
    }
    do {
      $item = floor($item / 3) - 2;
      $ret += max($item, 0);
    } while ($item > 0);
  }
  return $ret;
}

// 3249140
echo p1(get_input());

echo "\n";

// 4870838
echo p2(get_input());
