class Questionnaire::ReportsController < ApplicationController
  before_action :must_be_admin!, except: [:show]
  layout 'personal'

  def index
    @reports = initialize_grid(Questionnaire::Report.select(:id, :title, :updated_at).where(lang: params[:locale]).order(id: :desc))
  end

  def new
    @report = Questionnaire::Report.new
  end

  def create
    r = Questionnaire::Report.new __params
    r.lang = params[:locale]
    if r.save
      flash[:notice] = t 'messages.success'

      redirect_to questionnaire_reports_path
    else
      flash[:alert] = r.errors.full_messages
      @report = r
      render 'new'
    end
  end

  def show
    @report = Questionnaire::Report.find params[:id]
    render  layout:'cms'
  end

  def edit
    @report = Questionnaire::Report.find params[:id]
  end

  def update
    r = Questionnaire::Report.find params[:id]
    if r.update(__params)
      flash[:notice] = t 'messages.success'
      redirect_to questionnaire_reports_path
    else
      flash[:alert] = r.errors.full_messages
      @report = r
      render 'edit'
    end

  end

  def destroy
    Questionnaire::Report.destroy params[:id]
    flash[:notice] = t 'messages.success'
    redirect_to questionnaire_reports_path
  end

  private
  def __params
    params.require(:questionnaire_report).permit(:title, :summary)
  end
end