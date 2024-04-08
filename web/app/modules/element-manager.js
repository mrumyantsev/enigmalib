// An ElementManager is used to control interactive page elements,
// like: buttons, dropdown lists, etc.
export class ElementManager {

    #config;
    #elements;

    constructor(config) {
        this.#config = config;

        const rotorsContainers = [];
        const rotorsButtons = [];
        const positionsButtons = [];
        const ringsButtons = [];
        const rotorsLists = [];
        const positionsLists = [];
        const ringsLists = [];
    
        let rotorName;
        let positionName;
        let ringName;
    
        for (let i = 0; i < this.#config.maxRotors; i++) {
            rotorName = '#rotor' + (i+1);
            positionName = '#position' + (i+1);
            ringName = '#ring' + (i+1);
    
            rotorsContainers.push(new Elem(rotorName, 'inline'));
    
            rotorsButtons.push(new Elem(rotorName + '-button'));
            positionsButtons.push(new Elem(positionName + '-button'));
            ringsButtons.push(new Elem(ringName + '-button'));
    
            rotorsLists.push(new Elem(rotorName + '-list'));
            positionsLists.push(new Elem(positionName + '-list'));
            ringsLists.push(new Elem(ringName + '-list'));
        }
    
        this.#elements = {
            containers: {
                info: new Elem('#info'),
                machine: new Elem('#machine', 'inline'),
                reflector: new Elem('#reflector', 'inline'),
                rotors: rotorsContainers,
                plugboard: new Elem('#plugboard', 'inline'),
                encrypt: new Elem('#encrypt'),
                textFields: new Elem('.textfields', 'inline-flex')
            },
            buttons: {
                machine: new Elem('#machine-button'),
                reflector: new Elem('#reflector-button'),
                rotors: rotorsButtons,
                positions: positionsButtons,
                rings: ringsButtons,
                encrypt: new Elem('#encrypt-button')
            },
            lists: {
                machine: new Elem('#machine-list'),
                reflector: new Elem('#reflector-list'),
                rotors: rotorsLists,
                positions: positionsLists,
                rings: ringsLists
            },
            textareas: {
                plugboardtext: new Elem('#plugboardtext'),
                plaintext: new Elem('#plaintext'),
                ciphertext: new Elem('#ciphertext')
            },
            checkboxes: {
                groupBy5: new Elem('#groupBy5'),
                lowercase: new Elem('#lowercase')
            },
        };
    }

    get elements() {
        return this.#elements;
    }

    initElementsBySettings() {
        const machine = this.#config.machine;

        this.selectMachine(machine);

        this.#updateMachinesList();
    }

    selectMachine(idx) {
        this.#config.machine = idx;
    
        const machine = this.#config.spec.machines[this.#config.machine];
        const machineButton = this.#elements.buttons.machine;
        const reflectorButton = this.#elements.buttons.reflector;
        const rotorsButtons = this.#elements.buttons.rotors;
        const positionsButtons = this.#elements.buttons.positions;
        const ringsButtons = this.#elements.buttons.rings;
        const plugboardTextarea = this.#elements.textareas.plugboardtext;
        const setting = this.#config.settings[this.#config.machine];
    
        machineButton.text(machine.name);
    
        let r;
    
        if (machine.reflectorsCount) {
            r = setting.reflector;
    
            const reflectorName = machine.reflectors[r].name;
    
            reflectorButton.text(reflectorName);
        }
    
        let rotorName, positionNumber, ringNumber;
    
        for (let i = 0; i < machine.rotorsCount; i++) {
            r = setting.rotors[i].rotor;
    
            rotorName = machine.rotors[r].name;
            positionNumber = setting.rotors[i].position+1;
            ringNumber = setting.rotors[i].ring+1;
    
            rotorsButtons[i].text(rotorName);
            positionsButtons[i].text(positionNumber);
            ringsButtons[i].text(ringNumber);
        }
    
        if (machine.plugboardsCount) {
            plugboardTextarea.val(setting.plugboard);
        }
    
        this.#updateReflectorsList();
        this.#updateRotorPosRingsLists();
    
        this.updateVisibilities();
    }
    
    selectReflector(idx) {
        const setting = this.#config.settings[this.#config.machine];
        const machine = this.#config.spec.machines[this.#config.machine];
        const reflectorButton = this.#elements.buttons.reflector;
    
        setting.reflector = idx;
    
        const reflectorName = machine.reflectors[idx].name;
    
        reflectorButton.text(reflectorName);
    }
    
    selectRotor(num, idx) {
        const setting = this.#config.settings[this.#config.machine];
        const machine = this.#config.spec.machines[this.#config.machine];
        const rotorsButtons = this.#elements.buttons.rotors;
    
        setting.rotors[num].rotor = idx;
    
        const rotorName = machine.rotors[idx].name;
    
        rotorsButtons[num].text(rotorName);
    }
    
    selectPosition(num, idx) {
        const setting = this.#config.settings[this.#config.machine];
        const positionsButtons = this.#elements.buttons.positions;
    
        setting.rotors[num].position = idx;
    
        positionsButtons[num].text(idx+1);
    }
    
    selectRing(num, idx) {
        const setting = this.#config.settings[this.#config.machine];
        const ringsButtons = this.#elements.buttons.rings;
    
        setting.rotors[num].ring = idx;
    
        ringsButtons[num].text(idx+1);
    }
    
    displayInfo(text) {
        this.#hideInteract();
    
        const infoContainer = this.#elements.containers.info;
    
        infoContainer.text(text);
        infoContainer.show();
    }
    
    updateVisibilities() {
        const machine = this.#config.spec.machines[this.#config.machine];
        const reflector = this.#elements.containers.reflector;
        const rotors = this.#elements.containers.rotors;
        const plugboard = this.#elements.containers.plugboard;
        
        if (machine.reflectorsCount) {
            reflector.show();
        } else {
            reflector.hide();
        }
    
        for (let i = 0; i < machine.rotorsCount; i++) {
            rotors[i].show();
        }
    
        for (let i = machine.rotorsCount; i < this.#config.maxRotors; i++) {
            rotors[i].hide();
        }
    
        if (machine.plugboardsCount) {
            plugboard.show();
        } else {
            plugboard.hide();
        }
    }

    #updateMachinesList() {
        const machineList = this.#elements.lists.machine;
        const machines = this.#config.spec.machines;
    
        machineList.clear();
    
        let machineName;
    
        for (let i = 0; i < machines.length; i++) {
            machineName = machines[i].name;
    
            machineList.addItem(
                    machineName, `this.selectMachine(${i})`);
        }
    }
    
    #updateReflectorsList() {
        const machine = this.#config.spec.machines[this.#config.machine];
    
        if (!machine.reflectorsCount) {
            return;
        }
    
        const reflectorList = this.#elements.lists.reflector;
    
        reflectorList.clear();
    
        let machineName;
    
        for (let i = 0; i < machine.reflectors.length; i++) {
            machineName = machine.reflectors[i].name;
    
            reflectorList.addItem(
                    machineName, `this.selectReflector(${i})`);
        }
    }
    
    #updateRotorPosRingsLists() {
        const machine = this.#config.spec.machines[this.#config.machine];
    
        if (!machine.rotorsCount) {
            return;
        }
    
        const rotorsLists = this.#elements.lists.rotors;
        const positionsLists = this.#elements.lists.positions;
        const ringsLists = this.#elements.lists.rings;
    
        let rotorName;
    
        for (let j = 0; j < machine.rotorsCount; j++) {
            rotorsLists[j].clear();
            positionsLists[j].clear();
            ringsLists[j].clear();
    
            for (let i = 0; i < machine.rotors.length; i++) {
                rotorName = machine.rotors[i].name;
    
                rotorsLists[j].addItem(
                        rotorName, `this.selectRotor(${j}, ${i})`);
            }
    
            for (let i = 0; i < machine.rotors[j].positions; i++) {
                positionsLists[j].addItem(
                        i+1, `this.selectPosition(${j}, ${i})`);
            }
    
            for (let i = 0; i < machine.rotors[j].ringPositions; i++) {
                ringsLists[j].addItem(
                        i+1, `this.selectRing(${j}, ${i})`);
            }
        }
    }

    #hideInteract() {
        const machineContainer = this.#elements.containers.machine;
        const reflectorContainer = this.#elements.containers.reflector;
        const rotorsContainers = this.#elements.containers.rotors;
        const plugboardContainer = this.#elements.containers.plugboard;
        const encryptContainer = this.#elements.containers.encrypt;
        const textFieldsContainer = this.#elements.containers.textFields;
        
        machineContainer.hide();
        reflectorContainer.hide();
    
        for (let i = 0; i < rotorsContainers.length; i++) {
            rotorsContainers[i].hide();
        }
    
        plugboardContainer.hide();
        encryptContainer.hide();
        textFieldsContainer.hide();
    }

}

// An Elem is a wrapper for jQuery selector.
class Elem {

    #selector;
    #displayType;

    constructor(selectorText, displayType) {
        this.#selector = $(selectorText);

        if (!displayType) {
            this.#displayType = 'block';
        } else {
            this.#displayType = displayType;
        }
    }

    text(text) {
        this.#selector.text(text);
    }

    val(value) {
        if (value === undefined) {
            return this.#selector.val();
        }

        this.#selector.val(value);
    }

    clear() {
        this.#selector.html('');
    }

    addItem(text, onclick) {
        this.#selector.append(`<li>
  <a class="dropdown-item" onclick="${onclick}">
    ${text}
  </a>
</li>
`);
    }

    hide() {
        this.#selector.css('display', 'none');
    }

    show() {
        this.#selector.css('display', this.#displayType);
    }

    on(event, callback) {
        this.#selector.on(event, callback);
    }

}
