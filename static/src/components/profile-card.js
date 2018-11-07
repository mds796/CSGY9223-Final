import { html, PolymerElement } from '@polymer/polymer';
import {when} from 'lit-html/directives/when';

import './user-card.js'

/**
 * @customElement
 * @polymer
 */
class ProfileCard extends PolymerElement {
    static get properties() {
        return {
            offline: {type: Boolean, value: false},
            token: {type: String},
            _user: {type: Object, computed: '_getUser(token)'},
            _canLogin: {type: Object, computed: '_showLogin(token, offline)'},
            _hasUser: {type: Object, computed: '_showUser(token)'},
        };
    }

    static get template() {
        return html`
         <style>
            :host {
                display: block;
            }

            [hidden] {
                display: none;
            }
                        
            user-card {
                cursor: pointer;
            }
        
            a {
                color: var(--app-header-text-color);
                text-decoration: none;
                line-height: 30px;
            }
        </style>
        
        <dom-if if="[[_canLogin]]">
            <template>
                <a href="#/login">Log In</a>        
            </template>
        </dom-if>
        
        <dom-if if="[[_hasUser]]">
            <template>       
                <form action="/logout" method="post" id="logout" hidden>
                    <input type="submit" value="Log Out"/>       
                </form>
                <user-card title="Log out" user="[[_user]]" on-click="_userProfileClicked">
                </user-card>        
            </template>
        </dom-if>

        <dom-if if="[[offline]]">
            <template>
                <span>(Offline)</span>    
            </template>
        </dom-if>
    `;
    }

    _showLogin(token, offline) {
        return !(token || offline);
    }

    _showUser(token) {
        return token;
    }

    _getUser(token) {
        if (this.token) {
            // TODO: verify the token using the web service.
            return this.token.split(",")[0].split("=")[1];
        }
    }

    _userProfileClicked() {
        this._logOutClicked();
    }

    _logOutClicked() {
        this.shadowRoot.getElementById('logout').submit();
    }
}

window.customElements.define('profile-card', ProfileCard);
