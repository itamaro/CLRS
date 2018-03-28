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
