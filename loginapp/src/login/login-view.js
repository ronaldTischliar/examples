import { html, render } from "../../libs/lit-html.js";
import { fetchFromServer } from "./login-action.js"
import {store} from "../../store/store.js";

export default class LoginView extends HTMLElement {
    constructor() {
        super();
        this.user = "";
        this.password = "";
    }

    connectedCallback() {
        render(this.createView(), this);
    }

    createView() {
        return html`
            <form>
                <fieldset>
                    <a href="/info" id="next" hidden>Next</a>
                    <legend class="title">Login</legend>
                    <label class="label">User:
                        <input class="input is-primary" required id="user" name="user" placeholder="user"
                            @keyup=${e => this.onUserInput(e)} >
                    </label>
                    <label class="label">Password:
                        <input class="input is-primary" type="password" id="password" required name="password"
                            placeholder="password" minlength="8" @keyup=${e => this.onUserInput(e)} >
                    </label>
                    <button class="button is-primary is-fullwidth" @click="${e => this.login(e)}">Login</button>
            </form>
        `;
    }


    onUserInput({ target: { name, value } }) {
        this[name] = value;
    }

    login(event) {
        const { target: { form } } = event;
        event.preventDefault();
        form.reportValidity();
        if (form.checkValidity()) {
            console.log("end", this.user, this.password);
            fetchFromServer();
            document.getElementById("next") .click();
            store["user"]=this.user;
            store["token"]="testtoken";
        }


    }


}
customElements.define('r-login', LoginView);