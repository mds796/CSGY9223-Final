import {html, LitElement} from '@polymer/lit-element';
import {setPassiveTouchGestures} from '@polymer/polymer/lib/utils/settings.js';
import {installOfflineWatcher} from 'pwa-helpers/network.js';
import {installRouter} from 'pwa-helpers/router.js';
import {updateMetadata} from 'pwa-helpers/metadata.js';

import '@polymer/app-layout/app-drawer/app-drawer.js';
import '@polymer/app-layout/app-header/app-header.js';
import '@polymer/app-layout/app-scroll-effects/effects/waterfall.js';
import '@polymer/app-layout/app-toolbar/app-toolbar.js';

import {AppStyle} from './app-style.js';
import {menuIcon} from './icons.js';
import './register-view.js';
import './login-view.js';
import './feed-view.js';
import './follow-view.js';
import './about-view.js';
import './not-found-view.js';
import './user-profile.js';

/**
 * @customElement
 * @polymer
 */
class McubedApp extends LitElement {
    constructor() {
        super();

        this._drawerOpened = false;

        if (document.cookie !== "") {
            this._loginToken = document.cookie;
        }

        // To force all event listeners for gestures to be passive.
        // See https://www.polymer-project.org/3.0/docs/devguide/settings#setting-passive-touch-gestures
        setPassiveTouchGestures(true);
    }

    static get properties() {
        return {
            appTitle: {type: String},
            _page: {type: String},
            _drawerOpened: {type: Boolean},
            _loginToken: {type: String},
            _offline: {type: Boolean}
        };
    }

    render() {
        return html`
            ${AppStyle}
            <app-header condenses reveals effects="waterfall">
                <app-toolbar class="toolbar-top">
                    <button class="menu-btn" title="Menu" @click="${this._menuButtonClicked}">${menuIcon}</button>
                    <div main-title>${this.appTitle}</div>
                    <user-profile class="user-profile" .offline="${this._offline}" .token="${this._loginToken}" @logged-out="${this._loggedOut}"></user-profile>
                </app-toolbar>
            
                <!-- This gets hidden on a small screen-->
                <nav class="toolbar-list">
                    <a ?hidden="${!this._loginToken}" ?selected="${this._page === 'feed'}" href="/feed">Feed</a>
                    <a ?hidden="${!this._loginToken}" ?selected="${this._page === 'follow'}" href="/follow">Follow</a>
                    <a ?selected="${this._page === 'about'}" href="/about">About Us</a>
                </nav>
            </app-header>
            
            <app-drawer .opened="${this._drawerOpened}" @opened-changed="${this._drawerOpenedChanged}">
                <nav class="drawer-list">
                    <a ?hidden="${!this._loginToken}" ?selected="${this._page === 'feed'}" href="/feed">Feed</a>
                    <a ?hidden="${!this._loginToken}" ?selected="${this._page === 'follow'}" href="/follow">Follow</a>
                    <a ?selected="${this._page === 'about'}" href="/about">About Us</a>
                </nav>
            </app-drawer>

            <main role="main" class="main-content">
                <register-view class="page" ?active="${this._page === 'register'}"></register-view>
                <login-view class="page" ?active="${this._page === 'login'}" @logged-in="${this._loggedIn}"></login-view>
                
                <feed-view class="page" ?active="${this._page === 'feed'}"></feed-view>
                <follow-view class="page" ?active="${this._page === 'follow'}"></follow-view>
                <about-view class="page" ?active="${this._page === 'about'}"></about-view>
                
                <not-found-view class="page" ?active="${this._page === 'not-found'}"></not-found-view>              
            </main>

            <footer>
                <p>Made with &hearts; by the M<sup>3</sup> team.</p>
            </footer>
    `;
    }

    firstUpdated() {
        installRouter((location) => this._locationChanged(location));
        installOfflineWatcher((offline) => this._offlineChanged(offline));
    }

    updated(changedProps) {
        if (changedProps.has('_page')) {
            const pageTitle = this.appTitle + ' - ' + this._page;
            updateMetadata({
                title: pageTitle,
                description: pageTitle
                // This object also takes an image property, that points to an img src.
            });
        }
    }

    _offlineChanged(offline) {
        this._offline = offline;
        this._locationChanged();
    }

    _locationChanged() {
        const noUserPages = ['login', 'register'];
        const online = !this._offline;
        const loggedIn = this._loginToken;

        let page = this._extractPage();

        if (online && loggedIn && noUserPages.includes(page)) {
            window.history.pushState({}, '', '/feed');
            page = 'feed';
        } else if (online && !loggedIn && !noUserPages.includes(page)) {
            window.history.pushState({}, '', '/');
            page = 'about';
        } else if (!online) {
            window.history.pushState({}, '', '/');
            page = 'about';
        }

        this._loadPage(page);
        // Any other info you might want to extract from the path (like page type),
        // you can do here.

        // Close the drawer - in case the *path* change came from a link in the drawer.
        this._updateDrawerState(false);
    }

    _extractPage() {
        const path = window.decodeURIComponent(window.location.pathname);

        if (path === '/') {
            return this._loginToken ? 'feed' : 'about';
        } else {
            return path.slice(1);
        }
    }

    _updateDrawerState(opened) {
        if (opened !== this._drawerOpened) {
            this._drawerOpened = opened;
        }
    }

    _loadPage(page) {
        switch (page) {
            case 'feed':
                import('./feed-view.js');
                break;
            case 'follow':
                import('./follow-view.js');
                break;
            case 'about':
                import('./about-view.js');
                break;
            case 'register':
                import('./register-view.js');
                break;
            case 'login':
                import('./login-view.js');
                break;
            default:
                page = 'not-found';
                import('./not-found-view.js');
        }

        this._page = page;
    }

    _menuButtonClicked() {
        this._updateDrawerState(true);
    }

    _drawerOpenedChanged(e) {
        this._updateDrawerState(e.target.opened);
    }

    _loggedOut() {
        this._setCookie();
        window.history.pushState({}, '', '/');
        this._locationChanged();
    }

    _loggedIn() {
        const oneDay = 24 * 60 * 60 * 1000;
        const d = new Date();

        d.setTime(d.getTime() + (oneDay));

        this._setCookie(d.toUTCString());

        window.history.pushState({}, '', '/feed');
        this._locationChanged();
    }

    _setCookie(time) {
        const expires = time ? time : "Thu, 01 Jan 1970 00:00:00 GMT";

        document.cookie = "username=" + "mds796" + ";expires=" + expires + ";path=/";

        if (time) {
            this._loginToken = document.cookie
        } else {
            this._loginToken = undefined;
        }
    }
}

window.customElements.define('mcubed-app', McubedApp);
