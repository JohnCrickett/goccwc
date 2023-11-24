# goccwc
Go solution to [Coding Challenges](https://codingchallenges.fyi/challenges/intro) first challenge: [build your own wc tool](https://codingchallenges.fyi/challenges/challenge-wc)

## Testing

### Step 1

```bash
goccwc % wc -c test.txt
  342190 test.txt
goccwc % go run . -c test.txt
342190  test.txt
```

### Step 2

```bash
goccwc % wc -l test.txt
    7145 test.txt
goccwc % go run . -l test.txt
7145    test.txt
```

### Step 3

```bash
% wc -w test.txt
   58164 test.txt
goccwc % go run . -w test.txt
58164   test.txt
```

With the addition of some unit tests, which can be run with:
```bash
goccwc % go test .
ok      ccwc
```

### Step 4
```bash
goccwc % wc -m test.txt
  339292 test.txt
goccwc % go run . -m test.txt
339292  test.txt
```

### Step 5
```bash
% wc test.txt
    7145   58164  342190 test.txt
goccwc % go run . test.txt
7145    58164   342190  test.txt
```

### Step 6 (Final Step)
```bash
goccwc % cat test.txt | wc -l
    7145
goccwc % cat test.txt | go run . -l
7145
```

### Testing on Big Files (Over 100 GB)
```bash
goccwc % seq 1 300000 | xargs -Inone cat test.txt | wc
 2143500000 17449200000 102657000000
goccwc % seq 1 300000 | xargs -Inone cat test.txt | go run .
2143500000      17449200000     102657000000
```
Both use < 3MB memory.