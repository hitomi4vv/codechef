read line
set -- $line
c=$(($1-$2))
if [ $(($c % 10)) -eq 9 ]; then
  echo $(($c-1))
else
  echo $(($c+1))
fi
