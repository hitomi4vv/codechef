<?php
$input_lines = fgets(STDIN);
$ab = explode(' ', trim($input_lines));
$c = $ab[0] - $ab[1];
if ($c % 10 == 9) {
  echo $c - 1;
} else {
  echo $c + 1;
}
?>
