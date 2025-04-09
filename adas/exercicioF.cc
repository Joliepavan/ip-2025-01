#include <bits/stdc++.h>

using namespace std;

int main() {
    string input, resultado;
    cin >> input;

    string vogais = "aoyeui";

    for (char c : input) {
        c = tolower(c);
        if (vogais.find(c) == string::npos) {
            resultado += '.';
            resultado += c;
        }
    }

    cout << resultado << endl;

    return 0;
}
