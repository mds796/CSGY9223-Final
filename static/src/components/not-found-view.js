import { html, PolymerElement } from '@polymer/polymer';

import { ViewStyle } from './view-style.js';

/**
 * @customElement
 * @polymer
 */
class NotFoundView extends PolymerElement {
  static get template() {
    return html`
      <dom-if if=[[active]]>
        <template>
          ${ViewStyle}

          <section>
            <h2>Oops! You hit a 404</h2>
            <p>The page you're looking for doesn't seem to exist. Head back
              <a href="/">home</a> and try again?
            </p>
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

window.customElements.define('not-found-view', NotFoundView);
