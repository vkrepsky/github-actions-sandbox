import os

def main():
    filename = "dummy.txt"
    with open(filename, "w") as f:
        f.write(":) \n")
        f.close()

if __name__ == "__main__":
    main()
