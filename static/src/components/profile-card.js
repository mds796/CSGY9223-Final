import {html, LitElement} from '@polymer/lit-element';
import {when} from 'lit-html/directives/when';

import './user-card.js'

/**
 * @customElement
 * @polymer
 */
class ProfileCard extends LitElement {
    static get properties() {
        return {
            offline: {type: Boolean},
            token: {type: String},
            _user: {type: Object},
        };
    }

    constructor() {
        super();

        this.offline = false;
        this._user = {isLoggedIn: false};
    }

    render() {
        const offlineMessage = html`${when(this.offline, () => html`<span>(Offline)</span>`, () => html``)}`;

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
        
        
        <a href="/login" ?hidden="${this._user.isLoggedIn || this.offline}">Log In</a>
        <user-card title="Log out" user="${this._user.name}" ?hidden="${!this._user.isLoggedIn}" @click="${this._userProfileClicked}">    
        </user-card>

        ${offlineMessage}
    `;
    }

    updated(changedProps) {
        if (changedProps.has('token')) {
            if (this.token) {
                // TODO: verify the token using the web service.
                this._user = {name: "mds796", isLoggedIn: true};
            } else {
                this._user = {isLoggedIn: false};
            }
        }
    }

    _userProfileClicked() {
        this._logOutClicked();
    }

    _logOutClicked() {
        this.dispatchEvent(new CustomEvent('logged-out'));
    }
}

window.customElements.define('profile-card', ProfileCard);
