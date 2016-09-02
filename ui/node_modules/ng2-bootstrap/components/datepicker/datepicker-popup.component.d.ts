import { Renderer, ViewContainerRef } from '@angular/core';
import { NgModel } from '@angular/forms';
import { ComponentsHelper } from '../utils/components-helper.service';
export declare class DatePickerPopupDirective {
    cd: NgModel;
    viewContainerRef: ViewContainerRef;
    renderer: Renderer;
    componentsHelper: ComponentsHelper;
    private _activeDate;
    private _isOpen;
    private placement;
    private popup;
    constructor(cd: NgModel, viewContainerRef: ViewContainerRef, renderer: Renderer, componentsHelper: ComponentsHelper);
    activeDate: Date;
    private isOpen;
    hide(cb: Function): void;
    private show(cb);
}
