import { ElementManager } from "/app/modules/element-manager.js";
import { EventListener } from "/app/modules/event-listener.js";
import * as http from "/app/modules/http.js";

// An Application is the root class to run the Cipher Machines web
// application.
export class Application {

    #config = {
        maxRotors: 4,
        isGroupBy5: false,
        isLowercase: false,
        plaintext: "",
        ciphertext: "",
        spec: null,
        machine: 0,
        settings: null,
    };

    #elementManager;
    #eventListener;

    constructor() {
        this.#elementManager = new ElementManager(this.#config);
        this.#eventListener = new EventListener(this.#config, this.#elementManager);
    }

    run() {
        const response = http.getMachinesSpec();

        if (response.errorThrown) {
            console.debug("could not get machines spec. error:", response);

            this.#elementManager.displayInfo("server is offline");
            
            return;
        }

        this.#config.spec = response.data;

        console.debug("spec:", this.#config.spec);

        if (!this.#config.spec.machines) {
            this.#elementManager.displayInfo("no machines available");

            return;
        }

        this.#initMachinesSettings();

        this.#setMachineLike("M3");

        console.debug("config:", this.#config);

        this.#elementManager.initElementsBySettings();

        this.#eventListener.listen();
    }

    #setMachineLike(namePart) {
        this.#config.spec.machines.find(
            (el) => {
                if (el.name.includes(namePart)) {
                    this.#config.machine = this.#config.spec.machines.indexOf(el);
                }
            }
        );
    }

    #initMachinesSettings() {
        const settings = [];
    
        let machine;
        let setting;
        let rotor;
    
        for (let j = 0; j < this.#config.spec.machines.length; j++) {
            machine = this.#config.spec.machines[j];
            
            setting = {
                machine: j
            };
    
            if (machine.reflectorsCount) {
                setting.reflector = 0;
            }
    
            if (machine.rotorsCount) {
                setting.rotors = [];
    
                for (let i = 0; i < machine.rotorsCount; i++) {
                    rotor = {
                        rotor: i,
                        position: 0,
                        ring: 0
                    };
    
                    setting.rotors.push(rotor);
                }
            }
    
            if (machine.plugboardsCount) {
                setting.plugboard = "";
            }
    
            settings.push(setting);

            this.#config.settings = settings;
        }
    }

}
