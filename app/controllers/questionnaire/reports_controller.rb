class Questionnaire::ReportsController < ApplicationController
  before_action :must_be_admin!, except: [:show, :answer]
  layout 'personal'

  def result
    
  end

  def answer
    if recaptcha?
      uid = SecureRandom.uuid
      r = Questionnaire::Report.select(:id).find params[:report_id]
      r.questions.select(:id).each do |q|
        Questionnaire::Answer.create question_id: q.id, content:params["f_#{q.id}".to_sym], uid:uid
      end
      flash[:notice] = t 'messages.success'
    else
      flash[:alert] = t 'messages.failed'
    end
    redirect_to questionnaire_report_path(params[:report_id])
  end

  def index
    @reports = initialize_grid(Questionnaire::Report.select(:id, :title, :updated_at).order(id: :desc))
  end

  def new
    @report = Questionnaire::Report.new
  end

  def create
    r = Questionnaire::Report.new __params
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
    render layout: 'cms'
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