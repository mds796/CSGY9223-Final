import {html, PolymerElement} from '@polymer/polymer/polymer-element.js';
import './user-card.js';

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
          background: #cccccc;
          border-radius: 25px;
          padding: 5px;
          margin-bottom: 5px;
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
    `;
  }
  static get properties() {
    return {
      user: {
        type: String,
        value: ""
      },
      text: {
        type: String,
        value: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque ultrices leo sollicitudin nisl facilisis imperdiet. Nam a pellentesque enim. Donec sollicitudin placerat semper. Nam non neque quam. Suspendisse nec mauris rutrum dolor accumsan pellentesque nec vel tortor. Interdum et malesuada fames ac ante ipsum primis in faucibus. Cras et quam viverra nunc vulputate euismod nec in nisi. In vehicula faucibus erat, id ullamcorper sapien. Maecenas eu tristique ligula, a tempus ipsum. Nam vel pretium sed.'
      }
    };
  }
}

window.customElements.define('post-card', PostCard);
