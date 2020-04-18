cd ./vegeta/

cat targets.http | vegeta attack -rate 0 -max-workers 150 -duration 1m | tee results.bin | vegeta report
