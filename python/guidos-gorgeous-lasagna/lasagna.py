"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language: https://en.wikipedia.org/wiki/Guido_van_Rossum
"""

# TODO: define the 'EXPECTED_BAKE_TIME' constant
EXPECTED_BAKE_TIME = 40
# TODO: consider defining the 'PREPARATION_TIME' constant
#       equal to the time it takes to prepare a single layer
PREPARATION_TIME = 2

# TODO: define the 'bake_time_remaining()' function
def bake_time_remaining(elapsed_bake_time):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """
    return EXPECTED_BAKE_TIME - elapsed_bake_time

# TODO: define the 'preparation_time_in_minutes()' function
#       and consider using 'PREPARATION_TIME' here
def preparation_time_in_minutes(layers):
    """Calculate the time needed to prepare the lasagna.

    :param layers: int - number of layers in the lasagna.
    :return: int - time needed to prepare the lasagna.

    Function that takes the number of layers in the lasagna as an argument and
    returns the time needed to prepare the lasagna based on the `PREPARATION_TIME`.
    """
    return layers * PREPARATION_TIME

# TODO: define the 'elapsed_time_in_minutes()' function
def elapsed_time_in_minutes(layers, elapsed_bake_time):
    """Calculate the time elapsed.

    :param layers: int - number of layers in the lasagna.
    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - time elapsed.

    Function that takes the number of layers in the lasagna and the time the
    lasagna has been in the oven as arguments and returns the time elapsed
    based on the `PREPARATION_TIME`.
    """
    return layers * PREPARATION_TIME + elapsed_bake_time