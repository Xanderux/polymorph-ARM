## polymorph-ARM

**polymorph-ARM** is a tool that generates polymorphic versions of ARM assembly files. While the generated code may differ syntactically, it preserves the exact same behavior as the original.

This project addresses the current lack of polymorphic engines for ARM architectures, unlike the more mature tooling available for Intel x86.

### 🛠️ Use cases

**Red team:**
- Evade static detection mechanisms targeting known ARM shellcodes or patterns.

**Blue team:**
- Understand the variety of instruction combinations that achieve the same result.
- Generate hundreds of variants of the same shellcode — useful for training machine learning models or testing detection rules.

---

## 📦 Features

- Generates functionally equivalent but syntactically different versions of ARM assembly.
- Compatible with common instructions like `mov`, `add`, `sub`, etc.
- CLI interface for automation or batch processing.

---

## 🧱 Build

Make sure you have Go installed, then simply run:

```bash
make
```

This will build the `polymorph-ARM` binary in the current directory.

---

## 🚀 Usage

```bash
./polymorph-ARM -i input.s -o output.s
```

**Flags:**

```
  -i string
    	ARM assembly source file
  -o string
    	ARM assembly output file
```

---

## 🙏 Acknowledgments

Special thanks to:

- [Syscall59](https://x.com/syscall59) for planting the seed of this idea by documenting his [polymorphic x86 engine](https://medium.com/syscall59/writing-a-polymorphic-engine-73ec56a2353e).
- [Azeria](https://x.com/Fox0x01) for her outstanding content on ARM assembly. Her book [*Blue Fox: Arm Assembly Internals and Reverse Engineering*](https://www.amazon.fr/Blue-Fox-Assembly-Internals-Analysis/dp/1119745306) was a constant companion during the development of this engine.
