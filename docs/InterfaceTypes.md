## 📱 Interface Types

I often use the following thought process when deciding whether to return or pass things as pointers vs. by value.

1. Sharing down typically stays on the stack.
2. Sharing up typically escapes to the heap.

For a more context, be sure to watch this [video](https://youtu.be/ZMZpH4yT7M0?si=7kz9rzjgbXUcMs1v&t=459).

In most cases, the structs being copied are also fairly small.