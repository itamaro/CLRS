# Copyright 2018 Itamar Ostricher

"""My Red/Black tree implementation (educational purposes)"""


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
      # inserting root
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
    pass  # TODO implement
