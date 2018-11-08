import {html, PolymerElement} from '@polymer/polymer/polymer-element.js';

/**
 * @customElement
 * @polymer
 */
class DataProvider extends PolymerElement {
    static get template() {
        return html`
      <style>
        :host {
          display: hidden;
        }
      </style>
    `;
    }

    static get properties() {
        return {
            user: {type: Object, value: {}},
            posts: {type: Array, value: []},
            follows: {type: Array, value: []}
        };
    }

    ready() {
        this.fetchFeed(this);
    }

    fetchFeed(provider) {
        fetch('/feed').then(response => {
            return response.json();
        }).then(data => {
            console.log(data);
            provider.posts = data.posts;
        }).catch(err => {
            console.log("Unable to fetch feed: ", err);
        });
    }
}

window.customElements.define('data-provider', DataProvider);
