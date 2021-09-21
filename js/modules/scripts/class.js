export class Person{
    constructor(name) {
        this.name = name;
    }

    sayHallo(name){
        console.info(`hello ${name}, may name in ${this.name}`)
    }

}