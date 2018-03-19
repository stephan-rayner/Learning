#include <iostream>

using namespace std;

int requestStoppingValue(void) {
    int stoppingValue;
    cout << "Please provide a stopping value: ";
    cin >> stoppingValue;
    cout << endl;
    return stoppingValue;
}

int main() {
    cout << "= Fizz Buzz Launched =" << endl;
    int stoppingValue = requestStoppingValue();
    for (int i=1; i <= stoppingValue; i++) {
        if (i % 3 == 0) {
            cout << "Fizz";
        }
        if (i % 5 == 0) {
            cout << "Buzz";
        }
        if (i % 5 && i % 3) {
            // This uses the fact that 0 is falsey.
            cout << i;
        }
        cout << endl;
    }

    return 0;
}