# winterm

A terminal holiday tree animation written in go and using tcell. Features falling snow, an ASCII evergreen tree, and customizable colored lights grouped together in twinkling triplets.

![winterm demo](assets/demo.gif)

---

## Features

- **Dynamic Snow Simulation:** Snowflakes drift and settle naturally down the terminal window.
- **Grouped Tree Lights:** Lights populate the tree in clusters of 3, synchronized by group color.
- **Color Themes:** Switch between different color palettes (neon, frozen, warm, berry, classic, all) via command-line arguments.
- **Responsive Layout:** Adjusts layout dynamically when resizing the terminal window.

---

## Installation

Ensure you have [Go](https://go.dev/) (1.21 or newer) installed.

### go install
```bash
    go install github.com/Tomassalgueiro/winterm@latest
```

### cloning and running
1. Clone the repository:
```bash
   git clone [https://github.com/Tomassalgueiro/winterm.git](https://github.com/Tomassalgueiro/winterm.git)
   cd winterm
   go install .
```

2. Execute the program:
```bash
    winterm
```

## Color themes
```bash
    # if you don't chose any theme it will just default to the classic
    winterm -theme <neon, frozen, warm, berry, classic, all>
```




