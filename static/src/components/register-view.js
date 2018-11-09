import { html, PolymerElement } from '@polymer/polymer';

import {ViewStyle} from './view-style.js';

/**
 * @customElement
 * @polymer
 */
class RegisterView extends PolymerElement {
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
            <h2>Register</h2>
            <dom-if if="[[error]]">
                <template>
                    <p>[[error]]</p>
                </template>
            </dom-if>    
           
            <form action="/register" method="post">
              <div>
                <label for="username">Username: </label>
                <input id="username" name="username" placeholder="Username"/>
              </div>
              <div>
                <label for="password">Password: </label>
                <input id="password" name="password" type="password" placeholder="Password"/>          
              </div>
              <div>
                <label for="password2">Re-enter Password: </label>
                <input id="password2" name="password2" type="password" placeholder="Password"/>          
              </div>
                
              <input type="submit" value="Register"/>
            </form>
          </section>
        </template>
      </dom-if>
    `;
    }

    static get properties() {
        return {
            active: {type: Boolean, value: false},
            error: {type: String, value: false}
        };
    }
}

window.customElements.define('register-view', RegisterView);
