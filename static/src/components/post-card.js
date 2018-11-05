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
        }

        user-card {
          display: inline;
        }
      </style>
      <div>
        <user-card user="[[user]]"></user-card>
        <p>[[text]]</p>
      </div>
    `;
  }
  static get properties() {
    return {
      user: {
        type: String,
        value: 'mds796'
      },
      text: {
        type: String,
        value: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque ultrices leo sollicitudin nisl facilisis imperdiet. Nam a pellentesque enim. Donec sollicitudin placerat semper. Nam non neque quam. Suspendisse nec mauris rutrum dolor accumsan pellentesque nec vel tortor. Interdum et malesuada fames ac ante ipsum primis in faucibus. Cras et quam viverra nunc vulputate euismod nec in nisi. In vehicula faucibus erat, id ullamcorper sapien. Maecenas eu tristique ligula, a tempus ipsum. Nam vel pretium sed.'
      }
    };
  }
}

window.customElements.define('post-card', PostCard);
