import { html } from '@polymer/polymer';

export const AppStyle = html`
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
    
    .error {
        text-align: center;
        background: #ff4e3d;
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
`;