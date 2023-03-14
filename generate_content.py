import os

def main():
    filename = "dummy.txt"
    with open(filename, "w") as f:
        f.write("this is a generated content\n")
        f.close()

if __name__ == "__main__":
    main()
