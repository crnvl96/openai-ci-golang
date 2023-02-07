def first_prime_above(num):
    def is_prime(n):
        if n <= 1:
            return False
        for i in range(2, int(n**0.5) + 1):
            if n % i == 0:
                return False
        return True

    num += 1
    while not is_prime(num):
        num += 1
    return num

print(first_prime_above(2456646))
