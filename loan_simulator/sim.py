import sys

total = int(sys.argv[1])
per_time = int(sys.argv[2])
year, month = (int(x) for x in sys.argv[3].split("/"))

while total > 0:
    print("year: %d month: %d total: %d" % (year, month, total))

    month += 1
    if month > 12:
        month = 1
        year += 1

    total -= per_time
    if total < per_time:
        # print("year: %d month: %d total: %d" % (year, month, total))
        break
