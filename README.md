# goccwc
Go solution to [Coding Challenges](https://codingchallenges.fyi/challenges/intro) first challenge: [build your own wc tool](https://codingchallenges.fyi/challenges/challenge-wc)

## Testing

### Step 1

```bash
goccwc % wc -c test.txt
  342190 test.txt
goccwc % go run ./cmd/ccwc -c test.txt
342190  test.txt
```

### Step 2

```bash
goccwc % wc -l test.txt
    7145 test.txt
goccwc % go run ./cmd/ccwc -l test.txt
7145    test.txt
```

### Step 3

```bash
% wc -w test.txt
   58164 test.txt
goccwc % go run ./cmd/ccwc -w test.txt
58164   test.txt
```

With the addition of some unit tests, which can be run with:
```bash
goccwc % go test ./cmd/ccwc
ok      ccwc/cmd/ccwc
```

### Step 4
```bash
goccwc % wc -m test.txt
  339292 test.txt
goccwc % go run ./cmd/ccwc -m test.txt
339292  test.txt
```