const urlGetMachinesSpec = "http://127.0.0.1:8080/machines-spec";
const urlPostPlaintext = "http://127.0.0.1:8080/encryption";

// getMachinesSpec gets all the machines spec from the server.
export const getMachinesSpec = () => {
    console.debug("getting machines spec...");

    return doHttpRequest(urlGetMachinesSpec);
}

// postPlaintext posts plaintext with the machine setting to server and
// receives ciphertext.
export const postPlaintext = (config) => {
    console.debug("posting plaintext...");

    const machine = config.spec.machines[config.machine];
    const setting = config.settings[config.machine];
    const machineName = machine.name;

    let reflectorName = null;

    if (machine.reflectorsCount) {
        reflectorName = machine.reflectors[setting.reflector].name;
    }

    let rotorSettings = null;

    if (machine.rotorsCount) {
        rotorSettings = [];

        let r;

        for (let i = 0; i < setting.rotors.length; i++) {
            r = setting.rotors[i].rotor;

            rotorSettings.push({
                name: machine.rotors[r].name,
                position: setting.rotors[i].position,
                ring: setting.rotors[i].ring
            });
        }
    }

    let plugboardText = null;

    if (machine.plugboardsCount) {
        plugboardText = setting.plugboard;
    }
    
    const data = {
        plaintext: config.plaintext,
        machine: machine.name,
        settings: {
            reflector: reflectorName,
            rotorSettings: rotorSettings,
            plugboard: plugboardText
        }
    }

    console.log("data out:", data);

    return doHttpRequest(urlPostPlaintext, "POST", JSON.stringify(data));
}

// doHttpRequest does a basic synchronous uncached HTTP request, using
// JSON data type as default.
const doHttpRequest = (url, method, data, dataType) => {
    let response = {};

    if (!dataType) {
        dataType = "json";
    }

    $.ajax({
        async: false,
        cache: false,
        url: url,
        type: method,
        data: data,
        dataType: dataType,
        success: (data, textStatus, jqXHR) => {
            response = {
                data: data,
                textStatus: textStatus,
                jqXHR: jqXHR
            };
        },
        error: (jqXHR, textStatus, errorThrown) => {
            response = {
                jqXHR: jqXHR,
                textStatus: textStatus,
                errorThrown: errorThrown
            };
        }
    });

    return response;
}
