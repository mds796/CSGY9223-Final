import {html, PolymerElement} from '@polymer/polymer';

import './user-card.js'

/**
 * @customElement
 * @polymer
 */
class ProfileCard extends PolymerElement {
    static get properties() {
        return {
            offline: {type: Boolean},
            token: {type: String},
            user: {type: Object},

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
                display: inline-block;
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
                <user-card title="Log out" user="[[user.name]]" on-click="_userProfileClicked">
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

    _userProfileClicked() {
        this._logOutClicked();
    }

    _logOutClicked() {
        this.shadowRoot.getElementById('logout').submit();
    }
}

window.customElements.define('profile-card', ProfileCard);
