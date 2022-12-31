## Get unit test coverage

```shell
go test -covermode=count  -v -cover  ./...   -coverprofile=cover.out &&
            go tool cover -func=cover.out | grep total | awk '{print substr($3, 1, length($3)-3)}' | {
                read -r message; EXPECTED_COV=80;
                if [ "$message" -lt $EXPECTED_COV ];
                then  echo "FAILED, as Coverage is $message lesser then $EXPECTED_COV";
                else  echo "PASSED, as Coverage is $message greater then $EXPECTED_COV";
                fi
            }
```