import * as http from "/app/modules/http.js";

// An EventListener is used for listening and reacting on events like:
// clicking buttons, input text, etc.
export class EventListener {

    #config;
    #elementManager;
    #elements;

    constructor(config, elementManager) {
        this.#config = config;
        this.#elementManager = elementManager;
        this.#elements = elementManager.elements;
    }

    listen() {
        const machineList = this.#elements.lists.machine;
        const reflectorList = this.#elements.lists.reflector;
        const rotorsLists = this.#elements.lists.rotors;
        const positionsLists = this.#elements.lists.positions;
        const ringsLists = this.#elements.lists.rings;
        const plugboardTextarea = this.#elements.textareas.plugboardtext;
        const plaintextTextarea = this.#elements.textareas.plaintext;
        const ciphertextTextarea = this.#elements.textareas.ciphertext;
        const encryptButton = this.#elements.buttons.encrypt;
        const groupBy5Checkbox = this.#elements.checkboxes.groupBy5;
        const lowercaseCheckbox = this.#elements.checkboxes.lowercase;
        const machine = this.#config.machine;
        const setting = this.#config.settings[machine];

        machineList.on("click", (event) => {
            const index = event.target.cmIndex;

            this.#elementManager.selectMachine(index);
        });

        reflectorList.on("click", (event) => {
            const index = event.target.cmIndex;

            this.#elementManager.selectReflector(index);
        });

        rotorsLists.forEach((list, number) => {
            list.on("click", (event) => {
                const index = event.target.cmIndex;

                this.#elementManager.selectRotor(number, index);
            });
        });

        positionsLists.forEach((list, number) => {
            list.on("click", (event) => {
                const index = event.target.cmIndex;

                this.#elementManager.selectPosition(number, index);
            });
        });

        ringsLists.forEach((list, number) => {
            list.on("click", (event) => {
                const index = event.target.cmIndex;

                this.#elementManager.selectRing(number, index);
            });
        });

        plugboardTextarea.on("input", () => {
            setting.plugboard = plugboardTextarea.val();
        });

        plaintextTextarea.on("input", () => {
            this.#config.plaintext = plaintextTextarea.val();
        });

        encryptButton.on("click", (event) => {
            const response = http.postPlaintext(this.#config);

            if (response.errorThrown) {
                console.debug("could not get response with ciphertext. error:", response.errorThrown);

                ciphertextTextarea.text("server error");

                return;
            }

            event.target.blur();

            this.#config.ciphertext = response.data.ciphertext;

            if (this.#config.isGroupBy5) {
                this.#config.ciphertext = this.#groupBy5(this.#config.ciphertext);
            }

            if (this.#config.isLowercase) {
                this.#config.ciphertext = this.#lowercase(this.#config.ciphertext);
            }

            ciphertextTextarea.text(this.#config.ciphertext);
        });

        groupBy5Checkbox.on("change", () => {
            this.#config.isGroupBy5 = !this.#config.isGroupBy5;

            if (this.#config.isGroupBy5) {
                this.#config.ciphertext = this.#groupBy5(this.#config.ciphertext);
            } else {
                this.#config.ciphertext = this.#removeSpaces(this.#config.ciphertext);
            }

            ciphertextTextarea.text(this.#config.ciphertext);
        });

        lowercaseCheckbox.on("change", () => {
            this.#config.isLowercase = !this.#config.isLowercase;

            if (this.#config.isLowercase) {
                this.#config.ciphertext = this.#lowercase(this.#config.ciphertext);
            } else {
                this.#config.ciphertext = this.#uppercase(this.#config.ciphertext);
            }

            ciphertextTextarea.text(this.#config.ciphertext);
        });
    }

    #groupBy5(text) {
        if (text === "") {
            return text;
        }

        return this.#removeSpaces(text).match(/.{1,5}/g).join(" ");
    }

    #removeSpaces(text) {
        if (text === "") {
            return text;
        }

        return text.replace(/ /g, "");
    }

    #lowercase(text) {
        if (text === "") {
            return text;
        }

        return text.toLowerCase();
    }

    #uppercase(text) {
        if (text === "") {
            return text;
        }
        
        return text.toUpperCase();
    }

}
