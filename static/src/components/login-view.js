import { html, PolymerElement } from '@polymer/polymer';

import {ViewStyle} from './view-style.js';

/**
 * @customElement
 * @polymer
 */
class LoginView extends PolymerElement {
  static get template() {
    return html`
      <dom-if if=[[active]]>
        <template>
          ${ViewStyle}

          <style>
            :host {
              text-align: center
            }
          </style>

          <section>
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
                        
            <button on-click="_register">Register</button>
          </section>
        </template>
      </dom-if>
    `;
  }
    
  static get properties() {
    return {
      active: {type: Boolean, value: false}
    };
  }
 
  _register() {
    this.dispatchEvent(new CustomEvent('logged-in'));
  }
}

window.customElements.define('login-view', LoginView);
