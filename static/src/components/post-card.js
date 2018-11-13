import {html, PolymerElement} from '@polymer/polymer/polymer-element.js';
import './user-card.js';
import './timestamp-card.js';

/**
 * @customElement
 * @polymer
 */
class PostCard extends PolymerElement {
  static get template() {
    return html`
      <style>
        :host {
          display: block;
        }

        user-card {
          display: inline;
          margin-right: 5px;
        }

        p {
          display: inline;
        }
      </style>

      <user-card user="[[user]]"></user-card>
      <p>[[text]]</p>
      <timestamp-card timestamp="[[timestamp]]"></timestamp-card>
    `;
  }
  static get properties() {
    return {
      user: {
        type: String,
        value: ""
      },
      timestamp: {
        type: String,
        value: ""
      },
      text: {
        type: String,
        value: ""
      }
    };
  }
}

window.customElements.define('post-card', PostCard);
