import { html,render } from "../../libs/lit-html.js";
import {store} from "../../store/store.js";

export default class InfoView extends HTMLElement { 
    constructor() { 
        super();
        
    }

    connectedCallback() { 
        console.log("store",store);
        render(this.createView(),this);
    }

    createView() { 
        return html`
            <h1>Welcome ${store["user"]}</h1>
        `;
    }     
        
   
}
customElements.define('r-info',InfoView);