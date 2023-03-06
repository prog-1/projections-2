Projections on an arbitrary plane

1. Copy your program for the task 2. in the [projections-1](https://github.com/prog-1/projections-1) repository. Extend the program with support for a moving/rotating camera.

   Hints:

    1. The camera can be moved/rotated using a keyboard, i.e. by calling `ebiten.IsKeyPressed(...)` function with `ebiten.KeyArrowUp`, `ebiten.KeyArrowDown`, `ebiten.KeyArrowLeft`, `ebiten.KeyArrowRight`, `ebiten.KeyA`, `ebiten.KeyS`, `ebiten.KeyD`, `ebiten.KeyW` provided as arguments.
    
    2. You may represent a camera using vectors $\vec{s}$, $\vec{e_1}$, $\vec{e_2}$, $\vec{e_3}$, where $(\vec{e_1}, \vec{e_2}, \vec{e_3})$ is a coordinate system associated with the camera; $\vec{s}$ is a position of the camera.
    
    3. Implement a function that projects coordinates $(x, y, z)$ of an arbitrary point in the "world" coordinate system to coordinates $(x', y', z')$ in the "camera" coordinate system, e.g. `func (g Game) CameraProject(p Point) Point`.
    
    4. Start implementation with adding support for the camera movement in a single axis (e.g. `Y`).
