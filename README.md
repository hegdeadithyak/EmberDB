# EmberDB: A B-Tree-Based Database

**EmberDB** is a persistent database implemented using a B-Tree structure, optimized for disk-based storage. This README explains the design principles, node structure, and operational constraints of EmberDB.

---

## Key Features
- **Persistent Storage**: Uses disk-based page numbers instead of in-memory pointers.
- **Fixed Page Size**: Each page (node) is 4096 bytes (4 KB).
- **Balanced Tree Structure**: Nodes split when they exceed the page size, maintaining B-Tree properties.

---

## Node Structure
Each node in the B-Tree is divided into four main parts:
1. **Header**: Metadata about the node.
2. **Key-Value Pairs**: The actual data stored in the node.
3. **Child Pointers**: Pointers to child nodes (in internal nodes).
4. **Offsets**: Locations of key-value pairs within the node.

### Node Size Formula
```plaintext
Node Size = Header + Key-Value Pairs + Child Pointers + Offsets
```

---

### 1. Header
The **header** is a fixed 4 bytes:
- **2 bytes**: Type of the node (`btype`), indicating whether the node is a leaf or internal node.
- **2 bytes**: Number of keys (`nkeys`) stored in the node.

---

### 2. Key-Value Pairs
Key-value pairs store the actual data in the node, represented as:
```plaintext
| klen | vlen | key | val |
|  2B  |  2B  | ... | ... |
```
- **`klen`** (2 bytes): Length of the key.
- **`vlen`** (2 bytes): Length of the value.
- **`key`**: Actual key data, of size `klen`.
- **`val`**: Actual value data, of size `vlen`.

The total size of key-value pairs depends on the number of keys and the sizes of their respective keys and values:
```plaintext
Size of Key-Value Pairs = Σ (4 bytes + klen + vlen) for each key-value pair
```

---

### 3. Child Pointers
Child pointers are only present in internal nodes. They are disk-based 64-bit addresses (8 bytes each):
- A node with `nkeys` has `nkeys + 1` child pointers.
```plaintext
Size of Child Pointers = (nkeys + 1) × 8 bytes
```

---

### 4. Offsets
Offsets indicate the location of key-value pairs within the node. Each offset is 2 bytes:
```plaintext
Size of Offsets = nkeys × 2 bytes
```

---

## Node Layout
The structure of a node is as follows:
| Section         | Size Calculation                          | Description                                     |
|------------------|------------------------------------------|------------------------------------------------|
| **Header**       | 4 bytes                                  | Metadata about the node.                       |
| **Child Pointers**| `(nkeys + 1) × 8 bytes`                 | Disk-based pointers to child nodes.            |
| **Offsets**      | `nkeys × 2 bytes`                        | Offsets to locate key-value pairs.             |
| **Key-Value Pairs**| `Σ (4 bytes + klen + vlen)`            | The actual data stored in the node.            |

---

## Splitting Nodes
When a node exceeds the page size of **4096 bytes**, it is split into two nodes:
1. The **middle key** is promoted to the parent node.
2. The remaining keys are divided into two new child nodes.
3. This ensures the B-Tree remains balanced and adheres to its constraints.

---

## Example
Assume:
- Page size = 4096 bytes.
- Node contains:
  - `nkeys = 5`
  - Key sizes: 10, 15, 12, 14, 18 bytes.
  - Value sizes: 50, 60, 55, 65, 70 bytes.

### Calculations:
1. **Header**: `4 bytes`
2. **Child Pointers**: `(5 + 1) × 8 = 48 bytes`
3. **Offsets**: `5 × 2 = 10 bytes`
4. **Key-Value Pairs**:
   ```plaintext
   Σ (4 + klen + vlen) = (4+10+50) + (4+15+60) + ... = 368 bytes
   ```

**Total Node Size**:
```plaintext
4 + 48 + 10 + 368 = 430 bytes
```

---

## Constraints
- If the total size exceeds **4096 bytes**, the node will be split.
- The structure ensures efficient key-value storage and retrieval while maintaining disk optimization.

---

Feel free to extend the implementation or contact us for further details!
```