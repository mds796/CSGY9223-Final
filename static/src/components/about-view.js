import {html, PolymerElement} from '@polymer/polymer';

import {ViewStyle} from './view-style.js';

/**
 * @customElement
 * @polymer
 */
class AboutView extends PolymerElement {
  static get template() {
    return html`
      <dom-if if=[[active]]>
        <template>
          ${ViewStyle}

          <section>
            <h2>About Us</h2>
            <p>M <sup>3</sup> is a Twitter clone created by Miguel, Mel, and Matheus.</p>
          </section>
        </template>
      </dom-if>
    `;
  }

  static get properties() {
    return {
      active: {type: Boolean, value: false}
    };
  }
}

window.customElements.define('about-view', AboutView);
