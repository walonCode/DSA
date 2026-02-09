import math


def get_estimated_spread(audiences_followers: list[int]) -> float:
    if len(audiences_followers) == 0:
        return 0
    total: int = 0
    for value in audiences_followers:
        total += value
    return (total/len(audiences_followers)) * (len(audiences_followers) ** 1.2)


def get_follower_predictions(follower_count: int, influencer_type: str, num_months: int) -> int:
    type: dict[str, int] = {
        "fitness": 4,
        "cosmetics": 3
    }
    return follower_count * type[influencer_type] ** num_months


def get_influence_score(num_followers: int, average_engagement_percentage: float) -> float:
    return average_engagement_percentage * math.log(num_followers, 2)


def num_possible_order(num_post: int) -> int:
    if num_post == 0:
        return 0

    total: int = 1
    for i in range(1, num_post + 1):
        total *= i
    return total


def decayed_followerd(initail_followers: int, fraction_lost_daily: float, day: int) -> float:
    retention = 1 - fraction_lost_daily
    return initail_followers * (retention ** day)


def log_scale(data: list[int], base: int) -> list[float]:
    if len(data) == 0:
        return []

    answer: list[float] = []
    for value in data:
        value = math.ceil(math.log(value, base))
        answer.append(value)
    return answer


def average_followers(nums: list[int]) -> float:
    if len(nums) == 0:
        return 0

    total = 0
    for value in nums:
        total += value
    return total / len(nums)


average = average_followers([1, 2, 3, 4, 5, 6, 7])
print(f"The average is {average}")

new_scale = log_scale([1, 10, 100, 1000], 10)
print(f"The new scale is {new_scale}")

remaining = decayed_followerd(100, .5, 2)
print(f"The remaining followers are {remaining}")

value = get_estimated_spread([7, 4, 3, 100, 765, 2344, 1, 2, 32])
print(f"The spread is {math.ceil(value)}")


total = get_follower_predictions(10, "fitness", 1)
print(f"Total follower is {total}")

score = get_influence_score(40000, 0.3)
print(f"The score is {math.ceil(score)}")

possibilites = num_possible_order(6)
print(f"Number of possibilites is {possibilites}")
