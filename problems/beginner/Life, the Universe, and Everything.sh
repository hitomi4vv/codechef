while [ 1 ]
do
  read line
  if [ $line -eq 42 ]; then
    break;
  fi
  echo $line
done
