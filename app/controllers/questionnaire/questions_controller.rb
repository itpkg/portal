class Questionnaire::QuestionsController < ApplicationController
  layout 'personal'

  def new
    @question = Questionnaire::Question.new report_id: params[:report_id]
  end

  def create
    r = Questionnaire::Question.new __params
    if r.save
      flash[:notice] = t 'messages.success'
      redirect_to edit_questionnaire_report_path(r.report)
    else
      flash[:alert] = r.errors.full_messages
      @report = r
      render 'new'
    end
  end


  def edit
    @question = Questionnaire::Question.find params[:id]
  end

  def update
    q = Questionnaire::Question.find params[:id]
    if q.update(__params)
      flash[:notice] = t 'messages.success'
      redirect_to edit_questionnaire_report_path(q.report)
    else
      flash[:alert] = q.errors.full_messages
      @report = q
      render 'edit'
    end

  end

  def destroy
    q = Questionnaire::Question.find params[:id]
    q.destroy
    flash[:notice] = t 'messages.success'
    redirect_to edit_questionnaire_report_url(q.report)
  end

  private
  def __params
    params.require(:questionnaire_question).permit :name, :flag, :args, :report_id
  end
end