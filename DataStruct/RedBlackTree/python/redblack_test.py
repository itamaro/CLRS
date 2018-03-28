# Copyright 2018 Itamar Ostricher

"""My Red/Black Tree tests"""

from random import randint

from redblack import RedBlackTree


def test_insert_find():
  t = RedBlackTree()
  t.insert(1)
  t.insert(2)
  t.insert(3)
  t.insert(4)
  assert t.find(1) is not None
  assert t.find(2) is not None
  assert t.find(3) is not None
  assert t.find(4) is not None
  assert t.find(5) is None


def test_insertion_repairs():
  t = RedBlackTree()
  # case 1 repair (just root):
  t.insert(15)
  assert t.check_rb_invariants()
  # case 2 repair (red child of black node)
  t.insert(10)
  assert t.check_rb_invariants()
  # case 4 repair with right rotation and no inside-to-outside rotation substep
  # that requires root update
  t.insert(1)
  assert t.check_rb_invariants()
  # case 3 repair - recolor (trivial recursion)
  t.insert(37)
  assert t.check_rb_invariants()
  # case 4 repair with left rotation and no inside-to-outside rotation substep
  # that doesn't update root
  t.insert(45)
  assert t.check_rb_invariants()
  # case 3 repair with case 2 in recursion
  t.insert(21)
  assert t.check_rb_invariants()
  # another case 2 repair, making 2 internal subtrees on right side
  t.insert(41)
  assert t.check_rb_invariants()
  # case 4 repair with right rotation of internal subtree followed by left
  # rotation of grandparent, with no root update
  t.insert(17)
  assert t.check_rb_invariants()
  # case 4 repair with left rotation of internal subtree followed by right
  # rotation of grandparent, with no root update
  t.insert(43)
  assert t.check_rb_invariants()
  # case 3 repair with case 4 in the recursion, including right rotation of internal
  # subtree followed by left rotation of grandparent, including root update
  t.insert(30)
  assert t.check_rb_invariants()
