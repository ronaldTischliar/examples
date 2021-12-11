import { Router } from "./libs/vaadin-router.js";
import { appName } from "../../app.config.js";
import './login/login-view.js';
import "./login/info-view.js";


const outlet = document.querySelector('.view');
const router = new Router(outlet);
router.setRoutes([
  {path: '/',     component: 'r-login'},
  {path: '/info',  component: 'r-info'}
]);

console.log("router initialized", appName);


