import { html, PolymerElement } from '@polymer/polymer';

import { ViewStyle } from './view-style.js';

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
            <h2>About view!</h2>
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
