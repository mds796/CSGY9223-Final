import {html} from '@polymer/lit-element';

import {PageViewElement} from './page-view-element.js';
import {SharedStyles} from './shared-styles.js';

/**
 * @customElement
 * @polymer
 */
class LoginView extends PageViewElement {
    render() {
        return html`
          ${SharedStyles}
          <style>
            :host {
                text-align: center
            }
          </style>
          
          <h2>Log In</h2>
          
          <form action="/api/login" method="post">
            <div>
                <label for="username">Username: </label>
                <input id="username" name="username" placeholder="Username"/>
            </div>
            <div>
                <label for="password">Password: </label>
                <input id="password" name="password" type="password" placeholder="Password"/>          
            </div>
            
            <input type="submit" value="Log In"/>
          </form>
          
          <button @click="${this._register}">Register</button>
        `;
    }

    _register() {
        this.dispatchEvent(new CustomEvent('logged-in'));
    }
}

window.customElements.define('login-view', LoginView);
