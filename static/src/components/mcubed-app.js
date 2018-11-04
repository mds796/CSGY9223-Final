import {html, LitElement} from '@polymer/lit-element';
import {setPassiveTouchGestures} from '@polymer/polymer/lib/utils/settings.js';
import {installOfflineWatcher} from 'pwa-helpers/network.js';
import {installRouter} from 'pwa-helpers/router.js';
import {updateMetadata} from 'pwa-helpers/metadata.js';

import '@polymer/app-layout/app-drawer/app-drawer.js';
import '@polymer/app-layout/app-header/app-header.js';
import '@polymer/app-layout/app-scroll-effects/effects/waterfall.js';
import '@polymer/app-layout/app-toolbar/app-toolbar.js';

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
            <style>
                :host {
                    --app-drawer-width: 256px;
                    display: block;
                    --app-primary-color: #E91E63;
                    --app-secondary-color: #293237;
                    --app-dark-text-color: var(--app-secondary-color);
                    --app-light-text-color: white;
                    --app-section-even-color: #f7f7f7;
                    --app-section-odd-color: white;
                    --app-header-background-color: white;
                    --app-header-text-color: var(--app-dark-text-color);
                    --app-header-selected-color: var(--app-primary-color);
                    --app-drawer-background-color: var(--app-secondary-color);
                    --app-drawer-text-color: var(--app-light-text-color);
                    --app-drawer-selected-color: #78909C;
                }
               
            
                app-header {
                    position: fixed;
                    top: 0;
                    left: 0;
                    width: 100%;
                    text-align: center;
                    background-color: var(--app-header-background-color);
                    color: var(--app-header-text-color);
                    border-bottom: 1px solid #eee;
                }
            
                .toolbar-top {
                    background-color: var(--app-header-background-color);
                }
            
                [main-title] {
                    font-family: 'Pacifico';
                    text-transform: lowercase;
                    font-size: 30px;
                    /* In the narrow layout, the toolbar is offset by the width of the
                    drawer button, and the text looks not centered. Add a padding to
                    match that button */
                    padding-right: 44px;
                    text-decoration: none;
                    color: var(--app-header-text-color);
                }
                
                .toolbar-list {
                    display: none;
                }
            
                .toolbar-list > a {
                    display: inline-block;
                    color: var(--app-header-text-color);
                    text-decoration: none;
                    line-height: 30px;
                    padding: 4px 24px;
                }
            
                .toolbar-list > a[selected] {
                    color: var(--app-header-selected-color);
                    border-bottom: 4px solid var(--app-header-selected-color);
                }
            
                .menu-btn {
                    background: none;
                    border: none;
                    fill: var(--app-header-text-color);
                    cursor: pointer;
                    height: 44px;
                    width: 44px;
                }               
                
                .user-profile {
                    background: none;
                    border: none;
                    fill: var(--app-header-text-color);
                    height: 44px;                  
                }
            
                .drawer-list {
                    box-sizing: border-box;
                    width: 100%;
                    height: 100%;
                    padding: 24px;
                    background: var(--app-drawer-background-color);
                    position: relative;
                }
            
                .drawer-list > a {
                    display: block;
                    text-decoration: none;
                    color: var(--app-drawer-text-color);
                    line-height: 40px;
                    padding: 0 24px;
                }
            
                .drawer-list > a[selected] {
                    color: var(--app-drawer-selected-color);
                }
            
                /* Workaround for IE11 displaying <main> as inline */
                main {
                    display: block;
                }
            
                .main-content {
                    padding-top: 64px;
                    min-height: 100vh;
                }
            
                .page {
                    display: none;
                }
            
                .page[active] {
                    display: block;
                }
            
                footer {
                    padding: 24px;
                    background: var(--app-drawer-background-color);
                    color: var(--app-drawer-text-color);
                    text-align: center;
                }
                
                a[hidden] {
                    display: none;
                }
            </style>
            
            <!-- Header -->
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
            
            <!-- Drawer content -->
            <app-drawer .opened="${this._drawerOpened}" @opened-changed="${this._drawerOpenedChanged}">
                <nav class="drawer-list">
                    <a ?hidden="${!this._loginToken}" ?selected="${this._page === 'feed'}" href="/feed">Feed</a>
                    <a ?hidden="${!this._loginToken}" ?selected="${this._page === 'follow'}" href="/follow">Follow</a>
                    <a ?selected="${this._page === 'about'}" href="/about">About Us</a>
                </nav>
            </app-drawer>
            <!-- Main content -->
            <main role="main" class="main-content">
                <register-view class="page" ?active="${this._page === 'register'}"></register-view>
                <login-view class="page" ?active="${this._page === 'login'}" @logged-in="${this._loggedIn}"></login-view>
                
                <feed-view class="page" ?active="${this._page === 'feed'}"></feed-view>
                <follow-view class="page" ?active="${this._page === 'follow'}"></follow-view>
                <about-view class="page" ?active="${this._page === 'about'}"></about-view>
                
                <not-found-view class="page" ?active="${this._page === 'not-found'}"></not-found-view>              
            </main>
            <footer>
                <p>Made with &hearts; by the M3 team.</p>
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
    }

    _locationChanged() {
        const noUserPages = ['login', 'register'];
        let page = this._extractPage();

        if (this._loginToken && noUserPages.includes(page)) {
            window.history.pushState({}, '', '/feed');
            page = 'feed';
        } else if (!this._loginToken && !noUserPages.includes(page)) {
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
