from collections import deque


mm = {
    'a': {'b': 2, 'c': 3},
    'b': {},
    'c': {'d': 3},
    'd': {'e': 4, 'f': 5, 'g': 6},
    'g': {'h': 2},
}


def preorder(root, parent_children_relations: dict[str, dict[str, int]]):

    # create an empty stack
    stack = deque()

    # start from root node (set current node to root node)
    curr = root

    multiplier: int = 1  # parent-tree-adjusted data, pre-multiplied.

    # if current node is None and stack is also empty, we're done
    while stack or curr:
        # current node non-null means we are in depth digging mode
        # and building up stack
        if curr:
            child_data_tuples = deque(sorted(parent_children_relations.get(curr, {}).items(), reverse=True))
            stack.append((curr, multiplier, child_data_tuples))
            if len(child_data_tuples):
                # this is left-like traversal, or, use of first value from unbounded .right array of values
                curr, data = child_data_tuples.pop()
                multiplier *= data
            else:
                # this is bottom of sub-children tree
                # child_data_tuples will be empty
                curr = None

        else:
            # current node is None
            curr, multiplier, child_data_tuples = stack.pop()
            if len(child_data_tuples):
                # consume children first
                stack.append((curr, multiplier, child_data_tuples))
                # use of next value from unbounded .right array of values
                curr, data = child_data_tuples.pop()
                multiplier *= data
            else:
                print('====== CONSUMED ========', curr, multiplier)
                yield curr, multiplier
                curr = None
                # nulling out forces stack.pop() on next loop

l = list(preorder('a', mm))

print(l, sum([
    v
    for k, v in l
]))
