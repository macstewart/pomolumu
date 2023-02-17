# Pomolumu

A simple cli pomodoro timer

## Behaviour
- Starts a focus timer with the desired number of minutes
- After the timer has expired, a "break" will begin automatically
    - The breaks times follow the pattern 5m -> 5m -> 5m -> 15m, ad infinitum
- After the break, the next focus timer will queue up and pause until started. Press <space> to resume

### Building
```
git clone https://github.com/macstewart/pomolumu
cd pomolumu
go build
```
For global usage, move or link the resulting executable to a folder in your PATH

### Usage
```
./pomolumu [minutes]
```
[minutes] is an optional number of minutes for each focus timer. The default is 25.

#### hotkeys
- `<space>` to pause and resume the timer
- `r` to reset the timer back to its initial value
- `q` to exit

![run](./assets/.gif)
