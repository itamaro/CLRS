# Copyright 2018 Itamar Ostricher

"""My Red/Black Tree implementation (educational purposes)"""

from colorama import Fore, Style


class TreeNode:
  """Single tree node"""

  def __init__(self, val, parent=None, black=True, left=None, right=None):
    self.val = val
    self.parent = parent
    self.black = black
    self.left = left
    self.right = right

  @property
  def is_black(self):
    return self.black

  @property
  def is_red(self):
    return not self.black


class RedBlackTree:
  """Red/Black Tree: self-balancing binary search tree

  h - height of the tree
  N - number of nodes in the tree
  h = ~ logN
  """

  nil_node = TreeNode(None, None, True, None, None)

  def __init__(self):
    self.root = self.nil_node

  def find(self, val):
    """Look for `val` in the tree, returning the associated node.

    If `val` is not in the tree, return `None`.

    Runtime: O(h) = O(logN)
    """
    n = self.root
    while n != self.nil_node:
      if n.val == val:
        return n
      n = n.left if val < n.val else n.right
    return None

  def height(self):
    """Return the height of the tree (O(N))"""
    return self._height_recursive(self.root)

  def insert(self, val):
    """Insert a new node with value `val` to the tree, rebalancing if needed.

    Runtime: O(h) = O(logN)
    """
    p = None
    n = self.root
    while n != self.nil_node:
      p = n
      n = n.left if val <= n.val else n.right
    if p is None:
      # inserting root (AKA "case 1")
      new_node = TreeNode(val, None, True, self.nil_node, self.nil_node)
      self.root = new_node
    else:
      # inserting child node (initially red)
      new_node = TreeNode(val, p, False, self.nil_node, self.nil_node)
      if val <= p.val:
        p.left = new_node
      else:
        p.right = new_node

    self._insert_repair(new_node)
    return new_node

  def check_rb_invariants(self):
    """Return True is the RB invariants hold, False otherwise.

    Also prints the invalid invariants that are found.
    """

    def red_nodes_have_black_children(node):
      """Return True if all red nodes only black children, False otherwise."""
      if node == self.nil_node:
        return True
      if node.is_red:
        if node.left.is_red or node.right.is_red:
          print(f'{Fore.RED}Red node {node.val} has red child!{Style.RESET_ALL}')
          return False
      return (red_nodes_have_black_children(node.left) and
              red_nodes_have_black_children(node.right))

    def uniform_black_height(node):
      """Return the black height of all paths from `node` to all leaves.

      If not all paths has the same black-height, return -1.
      """
      if node == self.nil_node:
        return 0
      left_black_height = uniform_black_height(node.left)
      if left_black_height == -1:
        return -1
      right_black_height = uniform_black_height(node.right)
      if right_black_height == -1:
        return -1
      if left_black_height != right_black_height:
        print(f'{Fore.RED}Node {node.val} has left black height {left_black_height} '
              f'!= right black height {right_black_height}!{Style.RESET_ALL}')
        return -1
      return left_black_height + (1 if node.is_black else 0)

    return (self.root.is_black and
            red_nodes_have_black_children(self.root) and
            uniform_black_height(self.root) > 0)

  def _height_recursive(self, n):
    """Recursively find the height of subtree rooted at node `n`."""
    if n == self.nil_node:
      return 0
    return 1 + max(self._height_recursive(n.left),
                   self._height_recursive(n.right))

  def _insert_repair(self, node):
    """Make color adjustments and rotations after an insert to preserve
       the RB invariants.
    """
    parent = node.parent
    if parent is None:
      # repairing root - just need to set it to black
      node.black = True
      return
    if parent.is_black:
      # insert case 2 - nothing to do, all is well
      return
    # parent is red, so it's not root, so the grandparent exists
    grandparent = parent.parent
    uncle = grandparent.left if parent == grandparent.right else grandparent.right
    # note - uncle could be a nil_node (which is black)
    if uncle.is_red:
      # insert case 3 - parent & uncle are red, grandparent must be black
      # recolor parent and uncle to black, and grandparent to red
      parent.black = True
      uncle.black = True
      grandparent.black = False
      # recursively repair grandparent
      self._insert_repair(grandparent)
    else:
      # insert case 4 - parent is red, uncle is black, grandparent must be black
      # first, if the node is "on the inside subtree", need to rotate with its
      # parent so the node becomes the parent and the parent becomes the child
      # "on the outside subtree" - since both the node and the parent are red,
      # this doesn't change any black heights, but also doesn't resolve the issue
      # of having a red node as a child of another red node
      if node == grandparent.left.right:
        # left rotation
        grandparent.left = node
        parent.right = node.left
        if node.left != self.nil_node:
          node.left.parent = parent
        parent.parent = node
        node.parent = grandparent
        node.left = parent
        node, parent = parent, node
      elif node == grandparent.right.left:
        # right rotation
        grandparent.right = node
        parent.left = node.right
        if node.right != self.nil_node:
          node.right.parent = parent
        parent.parent = node
        node.parent = grandparent
        node.right = parent
        node, parent = parent, node
      # now the node is "on the outside subtree", so all that's left is
      # a simple rotation and recoloring of the parent & grandparent
      if node == parent.left:
        # right rotation on grandparent
        grandparent.left = parent.right
        if parent.right != self.nil_node:
          parent.right.parent = grandparent
        parent.right = grandparent
      else:
        # left rotation on grandparent
        grandparent.right = parent.left
        if parent.left != self.nil_node:
          parent.left.parent = grandparent
        parent.left = grandparent
      # common parent updates
      parent.parent = grandparent.parent
      grandparent.parent = parent
      greatgrandparent = parent.parent
      if greatgrandparent is None:
        self.root = parent
      else:
        if greatgrandparent.left == grandparent:
          greatgrandparent.left = parent
        else:
          greatgrandparent.right = parent
      # recoloring
      parent.black = True
      grandparent.black = False


def print_tree(t):
  """Print a ASCII art representation of the binary tree `t`."""
  if t is None or t.root is None or t.root == t.nil_node:
    return
  d = t.height()
  blocks = [[''] * (2**d - 1) for _ in range(1 + 2 * (d - 1))]

  def fill_level(t, node, level, offset):
    """Recursively fill in subtree rooted at `node` in the blocks matrix."""
    if node is None or node == t.nil_node:
      return
    blocks[level][offset] = (f'{Fore.YELLOW if node.is_black else Fore.RED}'
                             f'{node.val}{Style.RESET_ALL}')
    delta = 2 ** (d - level // 2 - 2)
    if node.left and node.left != t.nil_node:
      blocks[level + 1][offset - 1] = '/'
      fill_level(t, node.left, level + 2, offset - delta)
    if node.right and node.right != t.nil_node:
      blocks[level + 1][offset + 1] = '\\'
      fill_level(t, node.right, level + 2, offset + delta)

  fill_level(t, t.root, 0, 2**(d-1) - 1)
  # print the ASCII tree built in matrix
  for block in blocks:
    print(''.join(entry.center(3) for entry in block))
