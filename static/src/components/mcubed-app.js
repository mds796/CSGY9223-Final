import { html, PolymerElement } from '@polymer/polymer';
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
import './profile-card.js';

/**
 * @customElement
 * @polymer
 */
class McubedApp extends PolymerElement {
    static get template() {
        return html`
            ${AppStyle}          

            <app-header condenses reveals effects="waterfall">
                <app-toolbar class="toolbar-top">
                    <button class="menu-btn" title="Menu" on-click="_menuButtonClicked">${menuIcon}</button>
                    <div main-title>mCubed</div>
                    <profile-card class="user-profile" offline="[[_offline]]" user="[[user]]" token="[[token]]"></profile-card>
                </app-toolbar>
            
                <!-- This gets hidden on a small screen-->
                <nav class="toolbar-list">
                    <a hidden$="[[!_loggedIn]]" selected$="[[_isFeedView]]" href="#/feed">Feed</a>
                    <a hidden$="[[!_loggedIn]]" selected$="[[_isFollowView]]" href="#/follow">Follow</a>
                    <a selected$="[[_isAboutView]]" href="#/about">About Us</a>
                </nav>
            </app-header>
            
            <app-drawer opened="{{_drawerOpened}}">
                <nav class="drawer-list">
                    <a hidden$="[[!_loggedIn]]" selected$="[[_isFeedView]]" href="#/feed">Feed</a>
                    <a hidden$="[[!_loggedIn]]" selected$="[[_isFollowView]]" href="#/follow">Follow</a>
                    <a selected$="[[_isAboutView]]" href="#/about">About Us</a>
                </nav>
            </app-drawer>

            <main role="main" class="main-content">
                <dom-if if="[[error]]">
                    <template>
                        <section class="error">[[_errorMessage]]</section>
                    </template>
                </dom-if>    
                
                <register-view class="page" active$="[[_isRegisterView]]"></register-view>
                <login-view class="page" active$="[[_isLoginView]]"></login-view>
                
                <feed-view class="page" active$="[[_isFeedView]]"></feed-view>
                <follow-view class="page" active$="[[_isFollowView]]"></follow-view>
                <about-view class="page" active$="[[_isAboutView]]"></about-view>
                
                <not-found-view class="page" active$="[[_isNotFoundView]]"></not-found-view>              
            </main>

            <footer>
                <p>Made with &hearts; by the M<sup>3</sup> team.</p>
            </footer>
    `;
    }

    ready() {
        super.ready();

        // To force all event listeners for gestures to be passive.
        // See https://www.polymer-project.org/3.0/docs/devguide/settings#setting-passive-touch-gestures
        setPassiveTouchGestures(true);

        installRouter((location) => this._locationChanged(location));
        installOfflineWatcher((offline) => this._offlineChanged(offline));

        this.fetchFeed(this);
    }

    static get properties() {
        return {
            token: { type: String, computed: 'fetchToken(cookies)' },
            cookies: {type: Object, computed: 'parseCookie()'},
            user: {type: Object, computed: 'fetchUser(cookies)'},
            error: {type: String, computed: 'fetchError(cookies)'},

            _errorMessage: {type: String, computed: 'errorMessage(error)'},
            _page: {type: String, value: "about", observer: "_pageChanged"},
            _drawerOpened: {type: Boolean, value: false},
            _offline: {type: Boolean, value: false},
            _loggedIn: {type: Boolean, computed: "_isLoggedIn(token)"},

            // View predicates
            _isFeedView: {type: Boolean, computed: "_isActive(_page, 'feed')"},
            _isFollowView: {type: Boolean, computed: "_isActive(_page, 'follow')"},
            _isRegisterView: {type: Boolean, computed: "_isActive(_page, 'register')"},
            _isLoginView: {type: Boolean, computed: "_isActive(_page, 'login')"},
            _isAboutView: {type: Boolean, computed: "_isActive(_page, 'about')"},
            _isNotFoundView: {type: Boolean, computed: "_isActive(_page, 'not-found')"}
        };
    }

    errorMessage(error) {
        return JSON.parse(error);
    }

    parseCookie() {
        const cookie = {};
        document.cookie.split("; ").map(p=>p.split("=")).map(p=>cookie[p[0]]=p[1]);
        return cookie;
    }

    fetchUser(cookie) {
        return {name: cookie.username};
    }

    fetchToken(cookie) {
        return cookie[cookie.username] || "";
    }

    fetchError(cookie) {
        return cookie.error;
    }

    _isActive(page, expected) {
        return page === expected;
    }

    _pageChanged(newValue, oldValue) {
        const pageTitle = "mCubed - " + this._page;

        updateMetadata({
            title: pageTitle,
            description: pageTitle
            // This object also takes an image property, that points to an img src.
        });
    }

    _isLoggedIn(token) {
        return token !== "";
    }

    _offlineChanged(offline) {
        this._offline = offline;
        this._locationChanged(window.location);
    }

    _locationChanged(location) {
        const noUserPages = ['login', 'register'];
        const online = !this._offline;
        const loggedIn = this.token;

        let page = this._extractPage(location);

        if (online && loggedIn && noUserPages.includes(page)) {
            window.history.pushState({}, '', '#/feed');
            page = 'feed';
        } else if (online && !loggedIn && !noUserPages.includes(page)) {
            window.history.pushState({}, '', '#/');
            page = 'about';
        } else if (!online) {
            window.history.pushState({}, '', '#/');
            page = 'about';
        }

        this._loadPage(page);
        // Any other info you might want to extract from the path (like page type),
        // you can do here.

        // Close the drawer - in case the *path* change came from a link in the drawer.
        this._drawerOpened = false;
    }

    _extractPage(location) {
        const path = window.decodeURIComponent(location.hash);

        if (path === '#/' || path === '#' || path === '') {
            return this._loggedIn ? 'feed' : 'about';
        } else {
            return path.slice(2);
        }
    }

    _loadPage(page) {
        switch (page) {
            case 'feed':
                break;
            case 'follow':
                break;
            case 'about':
                break;
            case 'register':
                break;
            case 'login':
                break;
            default:
                page = 'not-found';
        }

        this._page = page;
    }

    _menuButtonClicked() {
        this._drawerOpened = true;
    }
}

window.customElements.define('mcubed-app', McubedApp);
